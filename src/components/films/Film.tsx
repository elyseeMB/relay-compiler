import { graphql, useFragment } from "react-relay";
import type { Film_item$key } from "./__generated__/Film_item.graphql.ts";

type Props = {
  film: Film_item$key;
};

const Film = ({ film }: Props) => {
  const data = useFragment(
    graphql`
      fragment Film_item on Film {
        id
        title
        director
        releaseDate
      }
    `,
    film
  );

  return (
    <div>
      <h3>{data.title}</h3>
      <p>Réalisateur: {data.director}</p>
      <p>Sortie: {data.releaseDate}</p>
    </div>
  );
};

export default Film;
