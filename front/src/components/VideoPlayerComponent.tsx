import { type FC } from 'react';
import { css } from '@emotion/react';
import ReactPlayer from 'react-player';

interface Props {
  url: string;
}

const playerStyles = css`
  width: 100% !important;
  height: auto !important;
`;

const VideoPlayerComponent: FC<Props> = ({ url }) => {
  return <ReactPlayer url={url} css={playerStyles} controls={true} />;
};

export default VideoPlayerComponent;
