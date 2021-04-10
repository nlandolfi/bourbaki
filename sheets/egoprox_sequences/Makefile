# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make egoprox_sequences.pdf
	make graph.pdf

egoprox_sequences.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > egoprox_sequences.tex

egoprox_sequences.pdf: ../../*.tex ../../trademark.pdf *.tex egoprox_sequences.tex
	# pdflatex --file-line-error -interaction=nonstopmode egoprox_sequences.tex
	../../latexrun egoprox_sequences.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: egoprox_sequences.pdf
	open egoprox_sequences.pdf

o:
	make open

terms: egoprox_sequences.tex
	bbk_sheet -mode ts egoprox_sequences.tex

spell:
	aspell -c egoprox_sequences.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f egoprox_sequences.*
	make reset
	bbk_sheet -mode mk > Makefile