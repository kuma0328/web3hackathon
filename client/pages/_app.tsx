import '../styles/globals.css';
import type { AppProps } from 'next/app';
import { QueryClient, QueryClientProvider } from 'react-query';
import { ReactQueryDevtools } from 'react-query/devtools';
import { ChainId, ThirdwebProvider, useAddress } from '@thirdweb-dev/react';
export default function App({ Component, pageProps }: AppProps) {
  const queryClient = new QueryClient();
  const activeChainId = ChainId.Goerli;
  return (
    <ThirdwebProvider desiredChainId={activeChainId}>
      <QueryClientProvider client={queryClient}>
        <Component {...pageProps} />
        <ReactQueryDevtools initialIsOpen={false} />
      </QueryClientProvider>
    </ThirdwebProvider>
  );
}
