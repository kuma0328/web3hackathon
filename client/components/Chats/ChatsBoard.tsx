import { Chat } from './Chat';

// どのタブを選んでいるのかはstateで管理する
// クリックで各チャット欄に移動する感じで
export const ChatsBoard = () => {
  return (
    <div className="relative">
      <div className="absolute top-0 left-0 w-full">
        <Chat chatTitle="レシピ" className=" ml-10 bg-lime-400">
          <div className="p-20">aa</div>
        </Chat>
        <Chat chatTitle="ランダム" className="left-1/3  bg-red-400">
          <div className="p-20">bb</div>
        </Chat>
        <Chat chatTitle="豆知識" className="left-2/3 -ml-10 bg-teal-400">
          <div className="p-20">cc</div>
        </Chat>
      </div>
    </div>
  );
};
