import React from 'react';
import { TextField } from '@mui/material';

export const Languages = [
    "english",
    "russian",
    "ukranian"
]

export function SelectLanguage({def="", onSelect=function(lang=""){}}) {
    return  <TextField select label="Language" defaultValue={Languages.find(el=>el===def)}
        SelectProps={{
          native: true,
        }}
        helperText="Desired language"
        onChange={onSelect}
    >
    {Languages.map((lang,i) => (
      <option key={i} value={lang}>{lang}</option>
    ))}
    </TextField>
  }