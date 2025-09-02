/**
 * @generated SignedSource<<d3709810865722b1807554c5f686eab4>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import { ConcreteRequest } from 'relay-runtime';
export type Role = "ADMIN" | "CUSTOMER";
export type NewUser = {
  email: string;
  fullName: string;
  role?: Role | null | undefined;
};
export type AppCreateUserMutation$variables = {
  input: NewUser;
};
export type AppCreateUserMutation$data = {
  readonly createUser: {
    readonly email: string;
    readonly fullName: string;
    readonly id: string;
    readonly role: Role;
  };
};
export type AppCreateUserMutation = {
  response: AppCreateUserMutation$data;
  variables: AppCreateUserMutation$variables;
};

const node: ConcreteRequest = (function(){
var v0 = [
  {
    "defaultValue": null,
    "kind": "LocalArgument",
    "name": "input"
  }
],
v1 = [
  {
    "alias": null,
    "args": [
      {
        "kind": "Variable",
        "name": "input",
        "variableName": "input"
      }
    ],
    "concreteType": "User",
    "kind": "LinkedField",
    "name": "createUser",
    "plural": false,
    "selections": [
      {
        "alias": null,
        "args": null,
        "kind": "ScalarField",
        "name": "id",
        "storageKey": null
      },
      {
        "alias": null,
        "args": null,
        "kind": "ScalarField",
        "name": "fullName",
        "storageKey": null
      },
      {
        "alias": null,
        "args": null,
        "kind": "ScalarField",
        "name": "email",
        "storageKey": null
      },
      {
        "alias": null,
        "args": null,
        "kind": "ScalarField",
        "name": "role",
        "storageKey": null
      }
    ],
    "storageKey": null
  }
];
return {
  "fragment": {
    "argumentDefinitions": (v0/*: any*/),
    "kind": "Fragment",
    "metadata": null,
    "name": "AppCreateUserMutation",
    "selections": (v1/*: any*/),
    "type": "Mutation",
    "abstractKey": null
  },
  "kind": "Request",
  "operation": {
    "argumentDefinitions": (v0/*: any*/),
    "kind": "Operation",
    "name": "AppCreateUserMutation",
    "selections": (v1/*: any*/)
  },
  "params": {
    "cacheID": "267d0ecccccade4b180e99db4c41a161",
    "id": null,
    "metadata": {},
    "name": "AppCreateUserMutation",
    "operationKind": "mutation",
    "text": "mutation AppCreateUserMutation(\n  $input: NewUser!\n) {\n  createUser(input: $input) {\n    id\n    fullName\n    email\n    role\n  }\n}\n"
  }
};
})();

(node as any).hash = "1849031862949941fc73074971442ec1";

export default node;
