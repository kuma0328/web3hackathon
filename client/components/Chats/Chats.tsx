import { useSelectingChatStore } from '../../stores/SelectingChatStore';
import { Chat } from './Chat';
import { Comments } from './Comments/Comments';

export const Chats = () => {
  const { recipe, random, tips, setRecipeChat, setRandomChat, setTipsChat } =
    useSelectingChatStore();
  return (
    <div className="relative my-10 md:m-0">
      <Chat
        chatTitle="レシピ"
        className="ml-10 bg-lime-400"
        onClick={setRecipeChat}
      >
        {recipe && <Comments />}
      </Chat>
      <Chat
        chatTitle="ランダム"
        className="left-1/3  bg-red-400"
        onClick={setRandomChat}
      >
        {random && <Comments />}
      </Chat>
      <Chat
        chatTitle="豆知識"
        className="left-2/3 -ml-10 bg-teal-400"
        onClick={setTipsChat}
      >
        {tips && <Comments />}
      </Chat>
    </div>
  );
};
