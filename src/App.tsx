import { useToast } from "./hooks/useToast.tsx";
import { useButton } from "./hooks/useButton.tsx";
import { useEffect } from "react";
import { useMessage } from "./hooks/useTest.tsx";

function App() {
  const { pushToast } = useToast();

  function onSubmit() {
    pushToast({
      title: "je suis le titre",
      content: "je suis le contenu",
      // duration: 2,
    });
  }

  function onSubmit2() {
    pushToast({
      title: "je suis le titre 2",
      content: "je suis le contenu 2",
      // duration: 2,
    });
  }

  return (
    <div className="container is-max-desktop p-5">
      <div className="is-flex is-justify-content-center is-align-items-center">
        <button onClick={onSubmit} className="button is-primary">
          Ajouter un nouveau toast
        </button>

        <button onClick={onSubmit2} className="button is-secondary">
          Ajouter un nouveau toast 2
        </button>
      </div>
    </div>
  );
}

export default App;
