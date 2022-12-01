# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make orthogonal_triangular_decomposition.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

orthogonal_triangular_decomposition.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > orthogonal_triangular_decomposition.tex

orthogonal_triangular_decomposition.pdf: ../../*.tex ../../trademark.pdf *.tex orthogonal_triangular_decomposition.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode orthogonal_triangular_decomposition.tex
	../../latexrun orthogonal_triangular_decomposition.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: orthogonal_triangular_decomposition.pdf
	open orthogonal_triangular_decomposition.pdf

o:
	make open

terms: orthogonal_triangular_decomposition.tex
	bbk_sheet -mode ts orthogonal_triangular_decomposition.tex

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
	rm -f orthogonal_triangular_decomposition.*
	make reset
	bbk_sheet -mode mk > Makefile