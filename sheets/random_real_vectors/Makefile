# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make random_real_vectors.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

random_real_vectors.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > random_real_vectors.tex

random_real_vectors.pdf: ../../*.tex ../../trademark.pdf *.tex random_real_vectors.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode random_real_vectors.tex
	../../latexrun random_real_vectors.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: random_real_vectors.pdf
	open random_real_vectors.pdf

o:
	make open

terms: random_real_vectors.tex
	bbk sheet -mode ts random_real_vectors.tex

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
	rm -f random_real_vectors.*
	make reset
	bbk sheet -mode mk > Makefile