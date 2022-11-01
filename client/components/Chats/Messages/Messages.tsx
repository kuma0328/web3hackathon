import React from 'react';
import { Message } from './Message';

export const Messages = () => {
  return (
    <div className="mt-10 flex flex-col justify-center overflow-y-auto">
      <div className="h-[50vh] lg:h-[80vh]">
        {[0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10].map(() => {
          return <Message />;
        })}
      </div>
    </div>
  );
};
