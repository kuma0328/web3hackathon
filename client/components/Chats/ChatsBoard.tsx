import { useSelectingChatStore } from '../../stores/SelectingChatStore';
import { Chat } from './Chat';
import { Messages } from './Messages/Messages';

// どのタブを選んでいるのかはstateで管理する
// クリックで各チャット欄に移動する感じで
export const ChatsBoard = () => {
  const { recipe, random, tips, setRecipeChat, setRandomChat, setTipsChat } =
    useSelectingChatStore();
  return (
    <div className="relative my-10 md:m-0">
      <Chat
        chatTitle="レシピ"
        className="ml-10 bg-lime-400"
        onClick={setRecipeChat}
      >
        {recipe && <Messages />}
      </Chat>
      <Chat
        chatTitle="ランダム"
        className="left-1/3  bg-red-400"
        onClick={setRandomChat}
      >
        {random && <Messages />}
      </Chat>
      <Chat
        chatTitle="豆知識"
        className="left-2/3 -ml-10 bg-teal-400"
        onClick={setTipsChat}
      >
        {tips && <Messages />}
      </Chat>
    </div>
  );
};
