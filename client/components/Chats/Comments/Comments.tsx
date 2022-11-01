import React from 'react';
import { CommentInputs } from './CommentInputs';
import { Comment } from './Comment';

export const Comments = () => {
  return (
    <div className="my-5 overflow-y-auto">
      <CommentInputs onClick={() => 1} />
      <div className="h-[50vh] lg:h-[40vh]">
        {[0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16].map(
          (idx) => {
            return <Comment key={idx} />;
          }
        )}
      </div>
    </div>
  );
};
