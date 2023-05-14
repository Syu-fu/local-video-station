import { type FC } from 'react';
import { css } from '@emotion/react';
import { CircularProgress, Box } from '@mui/material';

const overlayStyle = css`
  display: flex;
  justify-content: center;
  align-items: center;
  position: fixed;
  z-index: 9999;
  left: 0;
  top: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
`;

const ProgressBarComponent: FC = () => (
  <Box css={overlayStyle}>
    <CircularProgress />
  </Box>
);

export default ProgressBarComponent;
