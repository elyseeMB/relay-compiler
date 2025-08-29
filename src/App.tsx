import { graphql, useMutation } from "react-relay";

import { useToast } from "./hooks/useToast.tsx";

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

const handleCreateUser = async (data: Record<string, any>) => {
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

function App({ name }: { name: string }) {
  const { pushToast } = useToast();
  const [mutation, isLoading] = useMutation(userMutation);

  return (
    <div className="container is-max-desktop p-5">
      <div className="is-flex is-justify-content-center is-align-items-center"></div>
      <button
        onClick={() =>
          handleCreateUser({
            fullName: "janeDoeloe",
            password: "janeDoe@gmail.com",
          })
        }
      >
        Create User
      </button>
    </div>
  );
}

export default App;
