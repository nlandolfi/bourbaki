# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make affine_transformations_of_real_convex_sets.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

affine_transformations_of_real_convex_sets.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > affine_transformations_of_real_convex_sets.tex

affine_transformations_of_real_convex_sets.pdf: ../../*.tex ../../trademark.pdf *.tex affine_transformations_of_real_convex_sets.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode affine_transformations_of_real_convex_sets.tex
	../../latexrun affine_transformations_of_real_convex_sets.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: affine_transformations_of_real_convex_sets.pdf
	open affine_transformations_of_real_convex_sets.pdf

o:
	make open

terms: affine_transformations_of_real_convex_sets.tex
	bbk_sheet -mode ts affine_transformations_of_real_convex_sets.tex

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
	rm -f affine_transformations_of_real_convex_sets.*
	make reset
	bbk_sheet -mode mk > Makefile