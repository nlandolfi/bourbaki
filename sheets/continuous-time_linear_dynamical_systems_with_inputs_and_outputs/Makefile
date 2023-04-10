# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make continuous-time_linear_dynamical_systems_with_inputs_and_outputs.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

continuous-time_linear_dynamical_systems_with_inputs_and_outputs.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > continuous-time_linear_dynamical_systems_with_inputs_and_outputs.tex

continuous-time_linear_dynamical_systems_with_inputs_and_outputs.pdf: ../../*.tex ../../trademark.pdf *.tex continuous-time_linear_dynamical_systems_with_inputs_and_outputs.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode continuous-time_linear_dynamical_systems_with_inputs_and_outputs.tex
	../../latexrun continuous-time_linear_dynamical_systems_with_inputs_and_outputs.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: continuous-time_linear_dynamical_systems_with_inputs_and_outputs.pdf
	open continuous-time_linear_dynamical_systems_with_inputs_and_outputs.pdf

o:
	make open

terms: continuous-time_linear_dynamical_systems_with_inputs_and_outputs.tex
	bbk_sheet -mode ts continuous-time_linear_dynamical_systems_with_inputs_and_outputs.tex

spell:
	aspell -c sheet.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f continuous-time_linear_dynamical_systems_with_inputs_and_outputs.*
	make reset
	bbk_sheet -mode mk > Makefile