# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make partially_directed_graphs.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

partially_directed_graphs.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > partially_directed_graphs.tex

partially_directed_graphs.pdf: ../../*.tex ../../trademark.pdf *.tex partially_directed_graphs.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode partially_directed_graphs.tex
	../../latexrun partially_directed_graphs.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: partially_directed_graphs.pdf
	open partially_directed_graphs.pdf

o:
	make open

terms: partially_directed_graphs.tex
	bbk_sheet -mode ts partially_directed_graphs.tex

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
	rm -f partially_directed_graphs.*
	make reset
	bbk_sheet -mode mk > Makefile