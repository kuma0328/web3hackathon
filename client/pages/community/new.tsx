import { Board } from '../../components/BaseParts/Board';
import { Button } from '../../components/BaseParts/Button';
import { CommunityInputs } from '../../components/Recipe/CommunityInputs';
import { RecipeInputs } from '../../components/Recipe/RecipeInputs';
import { Page } from '../../components/Wrapper/Page';

const recipe = () => {
  return (
    <Page wide={true}>
      <Board className="w-full py-5">
        <div className="grid grid-cols-1 lg:grid-cols-2">
          <CommunityInputs />
          <RecipeInputs />
        </div>
        <div className="flex justify-center">
          <Button onClick={() => console.log('submit')}>
            コミュニティを作成する
          </Button>
        </div>
      </Board>
    </Page>
  );
};

export default recipe;
