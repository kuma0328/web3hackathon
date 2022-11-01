import { useRouter } from 'next/router';
import { ChatsBoard } from '../../components/Chats/ChatsBoard';
import { Community } from '../../components/Community/Community';
import { Page } from '../../components/Wrapper/Page';
const communityDetail = () => {
  const router = useRouter();
  const { id } = router.query;
  return (
    <Page
      wide={false}
      className="flex min-h-screen flex-col items-center py-5 lg:px-5"
    >
      <div className="grid w-full grid-cols-1 gap-5 py-5 md:h-3/4 lg:grid-cols-2">
        <Community
          title="焼肉同好会"
          description="焼肉同好会では各地のさまざまな焼き肉を食べながら交流を深めて行きます"
          recipeLink="a"
          onClickJoin={() => 1}
        />
        <ChatsBoard />
      </div>
    </Page>
  );
};

export default communityDetail;
