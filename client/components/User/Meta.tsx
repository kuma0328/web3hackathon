import { ConnectWallet } from '@thirdweb-dev/react';
type TProps = {
  address: string;
};
export const Meta = ({ address }: TProps) => {
  return (
    <div className="my-5">
      <ConnectWallet />
    </div>
  );
};
