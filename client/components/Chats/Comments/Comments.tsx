import React from 'react';
import { CommentInputs } from './CommentInputs';
import { Comment } from './Comment';

export const Comments = () => {
  return (
    <div className="my-5 overflow-y-auto">
      <CommentInputs onClick={() => 1} />
      <div className="h-[50vh] lg:h-[40vh]">
        {[0].map((idx) => {
          return <Comment key={idx} />;
        })}
      </div>
    </div>
  );
};
