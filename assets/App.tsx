import React from "preact/compat";
import { RouterProvider } from "react-router";
import { router } from "./routes.tsx";

// const userMutation = graphql`
//   mutation AppCreateUserMutation($input: NewUser!) {
//     createUser(input: $input) {
//       id
//       fullName
//       email
//       role
//     }
//   }
// `;

function App({ name }: { name: string }) {
  return (
    <>
      <RouterProvider router={router} />
    </>
  );
}

export default App;
