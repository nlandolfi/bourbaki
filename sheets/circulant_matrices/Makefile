# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make circulant_matrices.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

circulant_matrices.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > circulant_matrices.tex

circulant_matrices.pdf: ../../*.tex ../../trademark.pdf *.tex circulant_matrices.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode circulant_matrices.tex
	../../latexrun circulant_matrices.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: circulant_matrices.pdf
	open circulant_matrices.pdf

o:
	make open

terms: circulant_matrices.tex
	bbk_sheet -mode ts circulant_matrices.tex

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
	rm -f circulant_matrices.*
	make reset
	bbk_sheet -mode mk > Makefile