# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make n-dimensional_space.pdf
	make graph.pdf

n-dimensional_space.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > n-dimensional_space.tex

n-dimensional_space.pdf: ../../*.tex ../../trademark.pdf *.tex n-dimensional_space.tex
	# pdflatex --file-line-error -interaction=nonstopmode n-dimensional_space.tex
	../../latexrun n-dimensional_space.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: n-dimensional_space.pdf
	open n-dimensional_space.pdf

o:
	make open

terms: n-dimensional_space.tex
	bbk_sheet -mode ts n-dimensional_space.tex

spell:
	aspell -c n-dimensional_space.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f n-dimensional_space.*
	make reset
	bbk_sheet -mode mk > Makefile