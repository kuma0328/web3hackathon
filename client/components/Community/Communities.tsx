import { Community } from './Community';

export const Communities = () => {
  return (
    <div className="grid grid-cols-1 gap-10 lg:grid-cols-3 xl:grid-cols-4">
      {[0, 1, 2, 3, 4, 5].map((idx) => {
        return (
          <Community
            title="焼肉同好会"
            description="焼肉同好会では各地のさまざまな焼き肉を食べながら交流を深めて行きます"
            link="/community/aaa"
            recipeLink="./"
            key={idx}
          />
        );
      })}
    </div>
  );
};
