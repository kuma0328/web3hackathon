import { Board } from '../../../components/BaseParts/Board';
import { Community } from '../../../components/Community/Community';
import { Recipe } from '../../../components/Recipe/Recipe';
import { Page } from '../../../components/Wrapper/Page';

const recipe = () => {
  return (
    <Page wide={true}>
      <Board className="w-full">
        <div className="grid grid-cols-1 lg:grid-cols-2">
          <Community
            title="焼肉同好会"
            description="焼肉同好会では各地のさまざまな焼き肉を食べながら交流を深めて行きます"
            onClickJoin={() => 1}
            className="sticky top-0 left-0 pt-10"
          />
          <Recipe />
        </div>
      </Board>
    </Page>
  );
};

export default recipe;
