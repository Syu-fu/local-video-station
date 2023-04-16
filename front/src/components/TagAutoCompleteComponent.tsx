import { type FC } from 'react';
import { TextField, Autocomplete } from '@mui/material';
import type Tag from '../types/tag';

interface TagAutocompleteProps {
  tags: Tag[];
  selectedTags: Tag[];
  onChange: (tags: Tag[]) => void;
}

const TagAutocompleteComponent: FC<TagAutocompleteProps> = ({
  tags,
  selectedTags,
  onChange,
}) => {
  return (
    <Autocomplete
      multiple
      fullWidth
      sx={{ marginTop: '8px', width: '100%' }}
      options={tags}
      value={selectedTags}
      getOptionLabel={(option) => option.name}
      filterOptions={(options, params) => {
        const filtered = options.filter(
          (option) =>
            option.name
              .toLowerCase()
              .includes(params.inputValue.toLowerCase()) ||
            option.nameReading
              .toLowerCase()
              .includes(params.inputValue.toLowerCase())
        );

        return filtered;
      }}
      filterSelectedOptions
      onChange={(_, data) => {
        onChange(data);
      }}
      renderInput={(params) => <TextField {...params} label="tag" />}
    />
  );
};

export default TagAutocompleteComponent;
