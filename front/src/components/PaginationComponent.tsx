import { type FC } from 'react';
import Pagination from '@mui/material/Pagination';
import { styled } from '@mui/material/styles';
import { useNavigate } from 'react-router-dom';

const StyledPagination = styled(Pagination)(({ theme }) => ({
  '& .MuiPaginationItem-root': {
    color: theme.palette.text.primary,
  },
  '& .Mui-selected': {
    backgroundColor: theme.palette.primary.main,
    color: theme.palette.primary.contrastText,
    '&:hover': {
      backgroundColor: theme.palette.primary.dark,
    },
  },
}));

type Props = {
  pageCount: number;
  currentPage: number;
};

const PaginationComponent: FC<Props> = ({ pageCount, currentPage }) => {
  const navigate = useNavigate();

  const handleChange = (_: React.ChangeEvent<unknown>, page: number) => {
    navigate(`?page=${page}`);
  };

  return (
    <StyledPagination
      count={pageCount}
      page={currentPage}
      onChange={handleChange}
      shape="rounded"
      size="large"
    />
  );
};

export default PaginationComponent;
