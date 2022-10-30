import { Community } from './Community';

export const Communities = () => {
  return (
    <div className=" grid grid-cols-5 gap-10">
      {[0, 1, 2, 3, 4, 5].map(() => {
        return (
          <Community
            title="焼肉同好会"
            description="焼肉同好会では各地のさまざまな焼き肉を食べながら交流を深めて行きます"
          />
        );
      })}
    </div>
  );
};
