import { useChatContentStore } from '../../../stores/ChatContentStore';
import { Button } from '../../BaseParts/Button';
import { AddressZero } from '@ethersproject/constants';
import sdk from '../../../web3/initialize-sdk';
type TProps = {
  onClick: () => void;
};
export const CommentInputs = ({ onClick }: TProps) => {
  const { chatContent, chatImage, setChatContent, setChatImage } =
    useChatContentStore();

  const tokenDeploy = (async () => {
    try {
      const tokenAddress = await sdk.deployer.deployToken({
        name: 'FOOToken',
        symbol: 'FOO',
        primary_sale_recipient: AddressZero,
      });
      console.log(
        '✅ Successfully deployed token module, address:',
        tokenAddress
      );
    } catch (error) {
      console.error('failed to deploy token module', error);
    }
  })();
  return (
    <div className="text-right">
      <div className=" m-3 rounded-md bg-slate-200 p-10">
        <input
          type="file"
          placeholder="画像を送付する"
          className="my-1 w-full rounded-md bg-white p-3"
          value={chatImage}
          onChange={(e) => setChatImage(e.target.value)}
        />
        <input
          type="text"
          className="w-full rounded-md p-2 pb-20"
          placeholder="テキスト"
          value={chatContent}
          onChange={(e) => setChatContent(e.target.value)}
        />
        <Button onClick={() => tokenDeploy} className="my-2">
          コメントする
        </Button>
      </div>
    </div>
  );
};
