# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make independent_set_of_real_vectors.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

independent_set_of_real_vectors.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > independent_set_of_real_vectors.tex

independent_set_of_real_vectors.pdf: ../../*.tex ../../trademark.pdf *.tex independent_set_of_real_vectors.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode independent_set_of_real_vectors.tex
	../../latexrun independent_set_of_real_vectors.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: independent_set_of_real_vectors.pdf
	open independent_set_of_real_vectors.pdf

o:
	make open

terms: independent_set_of_real_vectors.tex
	bbk_sheet -mode ts independent_set_of_real_vectors.tex

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
	rm -f independent_set_of_real_vectors.*
	make reset
	bbk_sheet -mode mk > Makefile