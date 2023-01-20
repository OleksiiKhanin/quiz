import React from 'react';
import { Grid, TextField } from "@mui/material"
import {SelectLanguage} from './select'

export function WordInput({word={lang: "", value: "", description: ""}, onInput=word=>{}}) {
    const setLang = el=>onInput({...word, lang: el.target.value})
    const setValue = el=>onInput({...word, value: el.target.value})
    const setDescription = el=>onInput({...word, description: el.target.value})
    return (
          <Grid container rowSpacing={2} columnSpacing={{xs:1, sm:0, md: 2}}>
            <Grid item sm={12} md={5} lg={2}>
              <SelectLanguage def={word.lang} onSelect={setLang}/>
            </Grid>
            <Grid item sm={12} md={6} lg={10}>
              <TextField onBlur={setValue} label="Word" fullWidth variant="outlined"/>
            </Grid>
            <Grid item xs={12}>
              <TextField onBlur={setDescription} multiline label="Description" fullWidth variant="outlined"/>
            </Grid>
        </Grid>
    )
  }
  