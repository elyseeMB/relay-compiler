import {
  Environment,
  FetchFunction,
  Network,
  RecordSource,
  Store,
} from "relay-runtime";
import { PropsWithChildren } from "preact/compat";
import { RelayEnvironmentProvider } from "react-relay";

export function buildEndpoint(path: string): string {
  const host = import.meta.env.VITE_API_URL;

  if (!host) {
    return path;
  }

  const formattedHost =
    host.startsWith("http://") || host.startsWith("http://")
      ? host
      : `http://${host}`;

  const url = new URL(formattedHost);

  if (path) {
    url.pathname = path.startsWith("/") ? path : `/${path}`;
  }

  return url.toString();
}

const source = new RecordSource();
const store = new Store(source, {
  queryCacheExpirationTime: 1 * 60 * 1000,
  gcReleaseBufferSize: 20,
});

const fetchRelay: FetchFunction = async (
  request,
  variables,
  _,
  uploadables
) => {
  console.log(uploadables, variables);

  const resp = await fetch("http://localhost:8080/api/console/v1/query", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Accept:
        "application/graphql-response+json; charset=utf-8, application/json; charset=utf-8",
    },
    body: JSON.stringify({
      operationName: request.name,
      query: request.text,
      variables,
    }),
  });
  if (!resp.ok) {
    throw new Error("Response failed.");
  }
  return await resp.json();
};

export const relayEnvironment = new Environment({
  network: Network.create(fetchRelay),
  store,
});

export function RelayProvider({ children }: PropsWithChildren) {
  return (
    <RelayEnvironmentProvider environment={relayEnvironment}>
      {children}
    </RelayEnvironmentProvider>
  );
}
