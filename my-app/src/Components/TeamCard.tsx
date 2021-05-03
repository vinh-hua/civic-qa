import './TeamCard.css';

export type TeamCardProps = {
    name: string,
    img: string,
    bio: string
}

export function TeamCard(props: TeamCardProps) {
    return(
        <div className="team-bio">
            <div className="team-name-picture">
                <h2>{props.name}</h2>
                <img className="teammate-img" src={props.img}></img>
            </div>
            <p className="bio">{props.bio}</p>
        </div>
    );
}