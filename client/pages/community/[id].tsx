import { useRouter } from 'next/router';
import { ChatsBoard } from '../../components/Chats/ChatsBoard';
import { Community } from '../../components/Community/Community';
import { Page } from '../../components/Wrapper/Page';
const communityDetail = () => {
  const router = useRouter();
  const { id } = router.query;
  return (
    <Page wide={true}>
      <div className="mx-5 grid w-full grid-cols-2 gap-5">
        <Community
          title="焼肉同好会"
          description="焼肉同好会では各地のさまざまな焼き肉を食べながら交流を深めて行きます"
        />
        <ChatsBoard />
      </div>
    </Page>
  );
};

export default communityDetail;
