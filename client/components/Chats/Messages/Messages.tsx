import React from 'react';
import { Inputs } from './Inputs';
import { Message } from './Message';

export const Messages = () => {
  return (
    <div className="my-5 mt-10 flex flex-col justify-center overflow-y-auto">
      <Inputs onClick={() => 1} />
      <div className="h-[50vh] lg:h-[40vh]">
        {[0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16].map(
          (idx) => {
            return <Message key={idx} />;
          }
        )}
      </div>
    </div>
  );
};
