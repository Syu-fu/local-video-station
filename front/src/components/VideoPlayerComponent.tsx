import { type FC } from 'react';
import { css } from '@emotion/react';
import ReactPlayer from 'react-player';

interface Props {
  url: string;
  thumbnail: string;
}

const playerStyles = css`
  width: 100% !important;
  height: auto !important;
`;

const VideoPlayerComponent: FC<Props> = ({ url, thumbnail }) => {
  return (
    <ReactPlayer
      url={url}
      css={playerStyles}
      light={<img src={thumbnail} alt="Thumbnail" width="100%" />}
      controls={true}
    />
  );
};

export default VideoPlayerComponent;
