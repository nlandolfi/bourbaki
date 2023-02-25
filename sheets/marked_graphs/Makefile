# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make marked_graphs.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

marked_graphs.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > marked_graphs.tex

marked_graphs.pdf: ../../*.tex ../../trademark.pdf *.tex marked_graphs.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode marked_graphs.tex
	../../latexrun marked_graphs.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: marked_graphs.pdf
	open marked_graphs.pdf

o:
	make open

terms: marked_graphs.tex
	bbk_sheet -mode ts marked_graphs.tex

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
	rm -f marked_graphs.*
	make reset
	bbk_sheet -mode mk > Makefile