open Core.Std
open Filename
open Unix

type fileinfo =
  {name  : string;
   path  : string;
   stats : Unix.stats;
  }

type filedata =
  {meta    : fileinfo;
   content : string option
  }
  
let get_fileinfo name =
  {name  = name;
   path  = Filename.realpath name;
   stats = Unix.stat name;}

    
