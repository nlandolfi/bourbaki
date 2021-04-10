# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make independent_identically_distributed.pdf
	make graph.pdf

independent_identically_distributed.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > independent_identically_distributed.tex

independent_identically_distributed.pdf: ../../*.tex ../../trademark.pdf *.tex independent_identically_distributed.tex
	# pdflatex --file-line-error -interaction=nonstopmode independent_identically_distributed.tex
	../../latexrun independent_identically_distributed.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: independent_identically_distributed.pdf
	open independent_identically_distributed.pdf

o:
	make open

terms: independent_identically_distributed.tex
	bbk_sheet -mode ts independent_identically_distributed.tex

spell:
	aspell -c independent_identically_distributed.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f independent_identically_distributed.*
	make reset
	bbk_sheet -mode mk > Makefile