/**
 * @generated SignedSource<<e6decf8a848096eddd31a1468f1181e6>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import { ReaderFragment } from 'relay-runtime';
import { FragmentRefs } from "relay-runtime";
export type Film_item$data = {
  readonly director: string | null | undefined;
  readonly id: string;
  readonly releaseDate: string | null | undefined;
  readonly title: string | null | undefined;
  readonly " $fragmentType": "Film_item";
};
export type Film_item$key = {
  readonly " $data"?: Film_item$data;
  readonly " $fragmentSpreads": FragmentRefs<"Film_item">;
};

const node: ReaderFragment = {
  "argumentDefinitions": [],
  "kind": "Fragment",
  "metadata": null,
  "name": "Film_item",
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
      "name": "title",
      "storageKey": null
    },
    {
      "alias": null,
      "args": null,
      "kind": "ScalarField",
      "name": "director",
      "storageKey": null
    },
    {
      "alias": null,
      "args": null,
      "kind": "ScalarField",
      "name": "releaseDate",
      "storageKey": null
    }
  ],
  "type": "Film",
  "abstractKey": null
};

(node as any).hash = "3862088c68d66148d58fbf066e0e941f";

export default node;
