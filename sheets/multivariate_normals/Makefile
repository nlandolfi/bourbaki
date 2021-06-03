# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make multivariate_normals.pdf

multivariate_normals.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > multivariate_normals.tex

multivariate_normals.pdf: ../../*.tex ../../trademark.pdf *.tex multivariate_normals.tex
	# pdflatex --file-line-error -interaction=nonstopmode multivariate_normals.tex
	../../latexrun multivariate_normals.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: multivariate_normals.pdf
	open multivariate_normals.pdf

o:
	make open

terms: multivariate_normals.tex
	bbk_sheet -mode ts multivariate_normals.tex

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
	rm -f multivariate_normals.*
	make reset
	bbk_sheet -mode mk > Makefile