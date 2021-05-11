# Deploying CivicQA To the Cloud

The basics on deploying CivicQA to a cloud platform.  
This document assumes you have basic knowledge of your cloud provider of choice, docker, and GitHub Actions.

## Prerequisites:
- An account with a cloud provider of your choice (We used DigitalOcean)
- An accounter with DockerHub
- Admin access to the CivicQA repository (Or a fork)

## Cloud Resources:
- 1 or more Ubuntu VMs with Docker installed
- 1 MySQL Cluster
- 1 Load Balancer, routing traffic to your VM(s)

# Steps:
### Create Cloud Resources:
Provision the resources listed above on your cloud provider, be sure to install docker on every VM.  
Create the following databases in your MySQL cluster (`defaultdb` will already be created on DigitalOcean):
- `defaultdb`
- `logdb`

Route traffic on port `:80` on your load balancer, to port `:80` on your VMs. If you will be using https/ssl, route traffic from `:443` on your load balancer, to `:80` on your VM.
### Initializing GitHub Secrets:

Navigate to your repositories secrets: `settings > secrets`.  
Add the following `repository secrets` values:
- `DB_DSN`: The connection string for your main MySQL Database (Format: `<*ADMIN USERNAME*>:<*PASSWORD*>@tcp(<*ADDRESS*>:<*PORT*>)/defaultdb?parseTime=true`)
- `DB_DSN_LOGS`: The connection string for your log MySQL Database (Format: `<*ADMIN USERNAME*>:<*PASSWORD*>@tcp(<*ADDRESS*>:<*PORT*>)/logdb?parseTime=true`)
- `DOCKER_TOKEN`: An access token for your docker account ([generate here](https://hub.docker.com/settings/security))
- `DOCKER_USER`: Your docker username.
- `SSH_HOST`: The IP Address of your "main" VM, this VM will be the `manager` node of your `docker swarm`.
- `SSH_PRIVATE_KEY`: The SSH private key of your "main" VM. This should be the same VM/Host as your `SSH_HOST`.

### Build Backend Container Images:
With the secrets above, you should be able to succesfully run the GitHub Actions workflow to build the backend docker container images. On the repository go to Actions (`Actions` tab). On the left, under `All workflows`, click `build backend`. Trigger the Workflow by clicking `Run workflow`, select the branch `main`, and click `Run workflow`. 

This Workflow checks out the code from `main`, logs into your docker account using the `DOCKER_USER` and `DOCKER_TOKEN` secrets, and runs the appropriate scripts to build and push the backend container images to dockerhub. 

### Initialize Docker Swarm Mode:
`Docker Swarm` is a way to run a collection of `docker services` across multiple hosts/VMs. ([more info](https://docs.docker.com/engine/swarm/))

`Docker Swarm` functions for us as a lighweight version of something like Kubernetes. It handles container orchestration over multiple nodes, durability in the event of a node failure, replication of services, load balancing, service discovery, and more. 

#### Initialize your manager node: 
- SSH into your main vm, and execute the following command:  
$ `docker swarm init`. (If your recieve an error about multiple IPs, re-run the command with the following the `--advertise-addr` option set to the VMs external IP).
- This command will produce output indicating that your VM has become the manager node of your swarm, you will also see a long generated token, **save this!!!**, this is your `SWARM TOKEN`. It functions as a password to let other VMs join your VM as nodes.
#### Initialize your other node(s) if you have any:
- For each additional VM you have, SSH in and run the following command:  
$ `docker swarm join --token <*SWARM TOKEN*> <*MAIN VM IP*>` 

Congratulations! You have now initialized your docker swarm, and are ready to deploy the backend.

### Deploy The Backend
Now we can use the `Deploy` GitHub Action Workflow to deploy our services to our docker swarm.

As before, navigate to the Actions tab of the repository, and this time run the `Deploy` workflow. This Workflow checks out the repository code, uses `scp` to copy the production `docker-compose` file onto the manager node, SSHs into the manager node, and runs the appropriate script to deploy our backend services accross the nodes of our docker swarm. The workflow finishes by logging the services to stdout. Note that the actual deployment may not finish for several seconds to a few minutes after this workflow completes, as the containers need time to spin up and establish connections with eachother. If you ssh into the manager and inspect the services, you may see that the containers restart several times during this process, this is normal. After several minutes, you should try running $ `docker service ls` on the manager node, ensure that the services are replicated accross the nodes.

### Tuning your deployment
The deployment and replication policies for all services can be tuned to fit your scaling needs in the production `docker-compose` file (In the repo: `services/docker-compose.prod.yml`). 


### DNS, SSL Termination, and More:
The following will require some knowledge about DNS, HTTPS, and your cloud provider, you may need to do some additional research to figure this out.  
Something to keep in mind is that changes to DNS and forwarding may not always be instant, and can take several minutes to hours to be reflected. When in doubt, try the following:
- Hard refresh in your browser (`CTRL+R`) to prevent seeing a cached version of your page.
- Try a different different device/browser/connection. Your router, ISP, or device may cache DNS results. Try using your phone (off WIFI of course).
- Just wait. DNS can take a long time to update, even longer than the TTL. Other features like HTTPS forwarding can depend on your cloud provider. Give it a couple hours if things don't work right away.

#### HTTPS/SSL Termination
CivicQA's services use HTTP to communicate with eachother. This is done intentionally over HTTPS. We assume that all traffic between the services can be trusted, and therefore we don't bother with managing certificates.

Instead we opt for the [SSL-termination](https://docs.digitalocean.com/products/networking/load-balancers/how-to/ssl-termination/) strategy, where we use encrypted HTTPS traffic while communicating with our end-user/client, but decrypt the traffic before it finally reaches our api gateway, but is safe within our cloud network. 

Your cloud providers Load Balancer/Application Gateway should have options to add SSL certs (you can generate these automatically with DigitalOcean). You should set up the following with your load balancer:
- Forwarding rules: `HTTPS on :443 -> HTTP :80` (using your cert)
- Redirect HTTP to HTTPS: `on`

If you are unable to do these things with your load balancer, you may need to try a different resource from your provider, some of our strategies require an [OSI level 7](https://en.wikipedia.org/wiki/List_of_network_protocols_(OSI_model)) load balancer rather than a level 4. The default load balaner on DigitalOcean will suffice, but you will need to do more research for other providers (I believe the Azure and AWS equivalents are Application Gateway and Application Load Balancer respectively).

#### DNS
Assuming you have domain, and your frontend running on GitHub pages, you will need to set up the following DNS entries for your domain. For our examples we will assume your domain is `civicqa.codes`, and your frontend is deployed at `team-ravl.github.io` and configured to use the `civicqa.codes` domain. (Note: if you are using HTTPS on your frontend, you **MUST** use it on your backend as well due to the mxied-content policy of GitHub Pages).

| Type | Hostname | Value |
| -----|----------|-------|
|  A   |api.civicqa.codes| <*load balancer ip*>|
|  A   |civicqa.codes|185.199.108.153| 
|  A   |civicqa.codes|185.199.109.153|
|  A   |civicqa.codes|185.199.110.153|
|  A   |civicqa.codes|185.199.111.153|
| CNAME| www.civicqa.codes|team-ravl.github.io|

Additionally, you will need `NS` records from `civicqa.codes` to your cloud providers nameservers, and your domain must be configured to point to these as well (depends on your domain provider, they should have documentation). 

The A records for the IPs starting with `185.199` are all for GitHub Pages nameservers, they will direct traffic from civicqa.codes and www.civicqa.codes to your frontend.


#### Deploying Updates
To update your backend, merge your changes to the `main` branch, and re-run the GitHub Actions `build backend` and `Deploy`. This will build and deploy updated versions of your backend containers, to your swarm. There is no need to re-create any cloud resources or your swarm. Be aware that userdata will be preserved, although login sessions may not as they are stored in Redis. Using a cloud managed Redis or some other persistent KV store may be a future option to maintain sessions between updates. Either way, your clients will still be able to log in to their existing accounts as normal. 








