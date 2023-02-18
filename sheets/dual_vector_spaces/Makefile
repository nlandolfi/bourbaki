# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make dual_vector_spaces.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

dual_vector_spaces.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > dual_vector_spaces.tex

dual_vector_spaces.pdf: ../../*.tex ../../trademark.pdf *.tex dual_vector_spaces.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode dual_vector_spaces.tex
	../../latexrun dual_vector_spaces.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: dual_vector_spaces.pdf
	open dual_vector_spaces.pdf

o:
	make open

terms: dual_vector_spaces.tex
	bbk_sheet -mode ts dual_vector_spaces.tex

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
	rm -f dual_vector_spaces.*
	make reset
	bbk_sheet -mode mk > Makefile