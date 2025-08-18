import { graphql, useLazyLoadQuery } from "react-relay";

import { useToast } from "./hooks/useToast.tsx";
import Film from "./components/films/Film.tsx";
import { AppQuery } from "./__generated__/AppQuery.graphql.ts";

function App({ name }: { name: string }) {
  const { pushToast } = useToast();

  const data = useLazyLoadQuery<AppQuery>(
    graphql`
      query AppQuery {
        allFilms {
          films {
            id
            ...Film_item
          }
        }
      }
    `,
    {}
  );

  return (
    <div className="container is-max-desktop p-5">
      <div className="is-flex is-justify-content-center is-align-items-center">
        {data.allFilms!.films!.map((film) => (
          <Film key={film!.id} film={film!} />
        ))}
      </div>
    </div>
  );
}

export default App;
