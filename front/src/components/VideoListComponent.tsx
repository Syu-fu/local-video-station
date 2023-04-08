import { type FC } from 'react';
import type Video from '../types/video';
import VideoListItemComponent from './VideoListItemComponent';

interface Props {
  videos: Video[];
}

const VideoListComponent: FC<Props> = ({ videos }) => {
  return (
    <>
      {videos.map((video) => (
        <VideoListItemComponent key={video.id} video={video} />
      ))}
    </>
  );
};

export default VideoListComponent;
