import {
  Environment,
  FetchFunction,
  Network,
  RecordSource,
  Store,
} from "relay-runtime";
import { PropsWithChildren } from "react";
import { RelayEnvironmentProvider } from "react-relay";

function buildEndpoint(path: string): string {
  const host = import.meta.env.VITE_API_URL;

  if (!host) {
    return path;
  }

  const formattedHost =
    host.startsWith("http://") || host.startsWith("https://")
      ? host
      : `https://${host}`;

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

const HTTP_ENDPOINT = "https://graphql.org/graphql/";

const fetchRelay: FetchFunction = async (request, variables) => {
  const resp = await fetch(HTTP_ENDPOINT, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ query: request.text, variables }),
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
