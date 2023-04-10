import { type FC } from 'react';
import ReactPlayer from 'react-player';

interface Props {
  url: string;
}

const VideoPlayerComponent: FC<Props> = ({ url }) => {
  return <ReactPlayer url={url} />;
};

export default VideoPlayerComponent;
