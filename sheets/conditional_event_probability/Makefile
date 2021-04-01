# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make conditional_event_probability.pdf
	make graph.pdf

conditional_event_probability.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > conditional_event_probability.tex

conditional_event_probability.pdf: ../../*.tex ../../trademark.pdf *.tex conditional_event_probability.tex
	pdflatex --file-line-error -interaction=nonstopmode conditional_event_probability.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: conditional_event_probability.pdf
	open conditional_event_probability.pdf

o:
	make open

terms: conditional_event_probability.tex
	bbk_sheet -mode ts conditional_event_probability.tex

spell:
	aspell -c conditional_event_probability.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f conditional_event_probability.*
	make reset
	bbk_sheet -mode mk > Makefile