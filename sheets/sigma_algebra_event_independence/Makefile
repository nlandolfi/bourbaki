# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make sigma_algebra_event_independence.pdf
	make graph.pdf

sigma_algebra_event_independence.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > sigma_algebra_event_independence.tex

sigma_algebra_event_independence.pdf: ../../*.tex ../../trademark.pdf *.tex sigma_algebra_event_independence.tex
	pdflatex --file-line-error -interaction=nonstopmode sigma_algebra_event_independence.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: sigma_algebra_event_independence.pdf
	open sigma_algebra_event_independence.pdf

o:
	make open

defs: sigma_algebra_event_independence.tex
	defs sigma_algebra_event_independence.tex

spell:
	aspell -c sigma_algebra_event_independence.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm sigma_algebra_event_independence.*
	make reset
	bbk_sheet -mode mk > Makefile