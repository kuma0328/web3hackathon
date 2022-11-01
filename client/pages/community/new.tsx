import { Board } from '../../components/BaseParts/Board';
import { CommunityInputs } from '../../components/Recipe/CommunityInputs';
import { RecipeInputs } from '../../components/Recipe/RecipeInputs';
import { Page } from '../../components/Wrapper/Page';

const recipe = () => {
  return (
    <Page wide={true}>
      <Board className="w-full">
        <div className="grid grid-cols-1 lg:grid-cols-2">
          <CommunityInputs />
          <RecipeInputs />
        </div>
      </Board>
    </Page>
  );
};

export default recipe;
