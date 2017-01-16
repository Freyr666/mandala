open Lwt
open Core.Std
open LTerm_geom

let get_files dir =
  let d = Sys.ls_dir dir in
  String.concat ~sep:"\n" d

let frame widget =
  let frame = new LTerm_widget.frame in
  frame#set widget;
  frame

let main () =
  let waiter, wakener = Lwt.wait () in

  let box = new LTerm_widget.hbox in

  let parent  = new LTerm_widget.label (get_files ".") in
  let current = new LTerm_widget.label (get_files "..") in

  let lframe = frame parent  in
  let rframe = frame current in

  box#add lframe;
  box#add rframe;

  Lazy.force LTerm.stdout
  >>= fun term ->
  LTerm.enable_mouse term
  >>= fun () ->
  Lwt.finalize
    (fun () -> LTerm_widget.run term ~save_state:false ~load_resources:false box waiter)
    (fun () -> LTerm.disable_mouse term)
  
let () = Lwt_main.run (main ())
