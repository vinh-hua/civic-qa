import './TeamCard.css';

export type TeamCardProps = {
    name: string,
    img: string,
    bio: string,
    linkedin: string
}

export function TeamCard(props: TeamCardProps) {
    return(
        <div className="team-bio">
            <div className="team-name-picture">
                <div className="pic-linkedin">
                    <h2>{props.name}</h2>
                    <a className="linkedin-btn" target="_blank" href={props.linkedin}><img className="linkedin-icon" src="./assets/icons/linkedin-logo.png"></img></a>
                </div>
                <img className="teammate-img" src={props.img}></img>
            </div>

            <p className="bio">{props.bio}</p>
        </div>
    );
}