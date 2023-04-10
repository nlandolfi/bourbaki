# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make mutually_independent_events.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

mutually_independent_events.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > mutually_independent_events.tex

mutually_independent_events.pdf: ../../*.tex ../../trademark.pdf *.tex mutually_independent_events.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode mutually_independent_events.tex
	../../latexrun mutually_independent_events.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: mutually_independent_events.pdf
	open mutually_independent_events.pdf

o:
	make open

terms: mutually_independent_events.tex
	bbk_sheet -mode ts mutually_independent_events.tex

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
	rm -f mutually_independent_events.*
	make reset
	bbk_sheet -mode mk > Makefile