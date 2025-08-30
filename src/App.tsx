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

  const handleCreateUser = async (data: Record<string, any>) => {
    console.log(data);
    const response = await fetch(
      "http://localhost:8080/api/console/v1/auth/register",
      {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        credentials: "include",
        body: JSON.stringify(data),
      }
    );

    if (response.ok) {
      console.log("ok");
    }
  };

  const handleSubmit: FormEventHandler<HTMLFormElement> = (e) => {
    e.preventDefault();
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
              <button
                onClick={() =>
                  handleCreateUser({
                    fullName: "johnDoe",
                    password: "janeDoe4@gmail.com",
                  })
                }
                className="button is-link"
              >
                Submit
              </button>
            </div>
          </div>
        </form>
      </div>
    </div>
  );
}

export default App;
