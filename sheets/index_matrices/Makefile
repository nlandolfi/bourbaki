# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make index_matrices.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

index_matrices.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > index_matrices.tex

index_matrices.pdf: ../../*.tex ../../trademark.pdf *.tex index_matrices.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode index_matrices.tex
	../../latexrun index_matrices.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: index_matrices.pdf
	open index_matrices.pdf

o:
	make open

terms: index_matrices.tex
	bbk_sheet -mode ts index_matrices.tex

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
	rm -f index_matrices.*
	make reset
	bbk_sheet -mode mk > Makefile