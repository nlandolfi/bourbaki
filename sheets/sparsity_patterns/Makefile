# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make sparsity_patterns.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

sparsity_patterns.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > sparsity_patterns.tex

sparsity_patterns.pdf: ../../*.tex ../../trademark.pdf *.tex sparsity_patterns.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode sparsity_patterns.tex
	../../latexrun sparsity_patterns.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: sparsity_patterns.pdf
	open sparsity_patterns.pdf

o:
	make open

terms: sparsity_patterns.tex
	bbk_sheet -mode ts sparsity_patterns.tex

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
	rm -f sparsity_patterns.*
	make reset
	bbk_sheet -mode mk > Makefile