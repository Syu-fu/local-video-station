import { type FC } from 'react';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import VideoListPage from '../pages/list';

const AppRouter: FC = () => {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/list" element={<VideoListPage />} />
      </Routes>
    </BrowserRouter>
  );
};

export default AppRouter;
