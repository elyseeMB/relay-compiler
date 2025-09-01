import { graphql, useMutation } from "react-relay";

import { useToast } from "./hooks/useToast.tsx";
import { Field } from "./components/ui/Field.tsx";
import React, { FormEvent, FormEventHandler } from "react";

const userMutation = graphql`
  mutation AppCreateUserMutation($input: NewUser!) {
    createUser(input: $input) {
      id
      fullName
      email
      role
    }
  }
`;

function App({ name }: { name: string }) {
  const { pushToast } = useToast();
  const [mutation, isLoading] = useMutation(userMutation);

  const handleSubmit: FormEventHandler<HTMLFormElement> = async (e) => {
    e.preventDefault();
    const formData = new FormData(e.currentTarget);

    const response = await fetch(
      "http://localhost:8080/api/console/v1/auth/register",
      {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        credentials: "include",
        body: JSON.stringify(Object.fromEntries(formData)),
      }
    );

    if (response.ok) {
      console.log("ok");
    }
  };

  return (
    <div className="stack">
      <div className="is-flex is-justify-content-center is-align-items-center">
        <form onSubmit={handleSubmit}>
          <h2 className="pb-4 is-size-5 color">Registration</h2>

          <Field
            htmlFor="fullname"
            placeholder="your fullname"
            name="fullname"
            label="Fullname"
          />

          <Field
            htmlFor="password"
            placeholder="your password"
            name="password"
            label="Password"
          />

          <div className="field">
            <div className="control">
              <button className="button is-link">Submit</button>
            </div>
          </div>
        </form>
      </div>
    </div>
  );
}

export default App;
