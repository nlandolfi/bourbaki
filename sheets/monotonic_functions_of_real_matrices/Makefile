# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make monotonic_functions_of_real_matrices.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

monotonic_functions_of_real_matrices.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > monotonic_functions_of_real_matrices.tex

monotonic_functions_of_real_matrices.pdf: ../../*.tex ../../trademark.pdf *.tex monotonic_functions_of_real_matrices.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode monotonic_functions_of_real_matrices.tex
	../../latexrun monotonic_functions_of_real_matrices.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: monotonic_functions_of_real_matrices.pdf
	open monotonic_functions_of_real_matrices.pdf

o:
	make open

terms: monotonic_functions_of_real_matrices.tex
	bbk_sheet -mode ts monotonic_functions_of_real_matrices.tex

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
	rm -f monotonic_functions_of_real_matrices.*
	make reset
	bbk_sheet -mode mk > Makefile